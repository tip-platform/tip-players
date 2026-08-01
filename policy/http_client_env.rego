package httpclientenv

# Environments que deben existir obligatoriamente
required_envs := {"local", "compose"}

# Campos obligatorios por environment
required_fields := {"health_host", "health_port", "grpc_host", "grpc_port"}

# 1. Verifica que existan los environments requeridos
deny contains msg if {
    some env in required_envs
    not input[env]
    msg := sprintf("falta el environment '%v'", [env])
}

# 2. Verifica que cada environment tenga todos los campos
deny contains msg if {
    some env, _ in input
    some field in required_fields
    not input[env][field]
    msg := sprintf("'%v' no tiene el campo '%v'", [env, field])
}

# 3. Puertos deben ser numéricos (aunque vengan como string)
deny contains msg if {
    some env, cfg in input
    some field in {"health_port", "grpc_port"}
    port := cfg[field]
    not regex.match(`^[0-9]+$`, port)
    msg := sprintf("'%v.%v' debe ser numérico, encontrado: '%v'", [env, field, port])
}

# 4. health_port y grpc_port no pueden ser iguales
deny contains msg if {
    some env, cfg in input
    cfg.health_port == cfg.grpc_port
    msg := sprintf("'%v': health_port y grpc_port no pueden ser el mismo puerto", [env])
}
