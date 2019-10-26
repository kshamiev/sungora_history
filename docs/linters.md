### Linter (линтеры)

библиотека: [golangci-lint](https://github.com/golangci/golangci-lint)

Запуск и проверка линтеров (конфигурация линтеров: `.golangci.yml`)

    golangci-lint run

Вывести справочник всех возможных и используемых линтеров:

    `golangci-lint help linters`
    `golangci-lint linters`

Вывести справку по командной строке:

    `golangci-lint run -h`

Для игнорирования линтеров используется коомментарий перед нужной строчкой, функцией, перед пакетом для файла

    // nolint[:lintName1,lintName2,...]
