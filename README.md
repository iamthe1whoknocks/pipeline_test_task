# pipeline_test_task
  Реализовать модель обработки данных в виде пайплайна, состоящего из следующих этапов
1. Подача на вход пакетов данных. Пакет данных = слайсу случайных целых чисел из 10 элементов. Новый пакет подается каждые N мс (N задается в виде env переменной)   
2. Обработка пакетов: нахождение 3-х наибольших чисел в пакете. Вход: слайс int из 10 элементов, выход: слайс из 3-х элементов. Обработка пакетов должна производиться M воркерами (M задается в виде env переменной)
3. Аккумулятор: суммирование чисел обработанных пакетов, полученных на предыдущем этапе, и запись в единую переменную int
4. Публикатор: вывод на консоль текущего значения аккумулятора каждые K секунд (K задается в виде env переменной)


# Manual

## Build and run service
`make up`

## Build binary
`make build`

## Run binary
`make run`

# Config
## config file 
`env.env`

## Values
1. `PIPELINE_NEW_INPUT_TIME_MS` - time (in ms) to send new data batch to handler 
2. `PIPELINE_WORKER_NUM` - handler workers count
3. `PIPELINE_PUBLICATOR_OUTPUT_TIME_SEC` - time (in sec) to publicate result of accumlator value in stdout


