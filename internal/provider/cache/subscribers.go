package cache

const subscribersKey = "subscribers"

func (p *Provider) AddSubscriber(address string) {
	subs := p.GetSubscribers()
	if len(subs) < 1 {
		p.subscribers.Set(subscribersKey, []string{
			address,
		})
		return
	}
	subs = append(subs, address)
	p.subscribers.Set(subscribersKey, subs)
}

func (p *Provider) GetSubscribers() []string {
	subs, ok := p.subscribers.Get(subscribersKey)
	if !ok {
		return []string{}
	}
	return subs.([]string)
}
