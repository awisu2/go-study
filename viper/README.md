# viper の勉强

github: [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)

- what is?: config, environment values management tool
- readable
  - files: JSON, TOML, YAML, HCL, envfile, java properties files.
    - it can watch each files.
  - "environment variables", "remote config systems"(watchiable), "command line flags" and "buffer"(ex: io.Buffer)
- precedence order (writed below)

## 課題

- jsonのメタ設定が効いていない感じがする
  - 変数名によって、大文字小文字の違いは吸収したが、メタでキー名を指定してもパラメータの取得ができなかった

## precedence order

each item takes precedence over the item below it.

- explicit call to `Set`
- flag
- env
- config
- key/value store
- default

