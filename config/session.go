package config


type
	// Session is  session's session interface, think it like a local store of each session's values
	// implemented by the internal session iteral, normally the end-user will never use this interface.
	// gettable by the provider -> sessions -> your app
	Session interface {
		ID() string
		Get(string) interface{}
		HasFlash() bool
		GetFlash(string) interface{}
		GetString(key string) string
		GetFlashString(string) string
		GetInt(key string) (int, error)
		GetInt64(key string) (int64, error)
		GetFloat32(key string) (float32, error)
		GetFloat64(key string) (float64, error)
		GetBoolean(key string) (bool, error)
		GetAll() map[string]interface{}
		GetFlashes() map[string]interface{}
		VisitAll(cb func(k string, v interface{}))
		Set(string, interface{})
		SetFlash(string, interface{})
		Delete(string)
		DeleteFlash(string)
		Clear()
		ClearFlashes()
	}

