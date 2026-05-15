# viper

Config management: YAML file, defaults, env var override.

## What's here

`main.go` — reads `config.yml`, sets a default, auto-reads `CRAWLER_*` env vars.  
`config.yml` — database URL, index URL, goroutine limit, network shares map.

## Todo

- [ ] Unmarshal into a typed struct with `viper.Unmarshal(&cfg)` instead of individual `GetString` calls — this is the idiomatic production pattern.
- [ ] Override `max_num_goroutine` via `CRAWLER_MAX_NUM_GOROUTINE=8` in the shell and verify Viper picks it up without changing the YAML.
- [ ] Write a test using a temp config file: point Viper at it, assert the parsed values — make config loading testable.
- [ ] Add `viper.WatchConfig` + `viper.OnConfigChange` to hot-reload a value at runtime without restarting the process.
