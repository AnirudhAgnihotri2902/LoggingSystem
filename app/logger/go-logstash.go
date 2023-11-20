package logger

func (l *Logstash) Log(payload map[string]interface{}) {
	l.pushJsonMessage(payload)
}

func (l *Logstash) Info(payload map[string]interface{}) {
	payload["level"] = "INFO"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Debug(payload map[string]interface{}) {
	payload["level"] = "DEBUG"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Warn(payload map[string]interface{}) {
	payload["level"] = "WARN"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Error(payload map[string]interface{}) {
	payload["level"] = "ERROR"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) LogString(stringMsg string) {
	l.pushStringMessage(stringMsg)
}
