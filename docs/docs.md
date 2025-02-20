# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [protos/proto/Service/Service.proto](#protos_proto_Service_Service-proto)
    - [CreateTaskRequest](#service-CreateTaskRequest)
    - [CreateTaskResponse](#service-CreateTaskResponse)
    - [DeleteTaskRequest](#service-DeleteTaskRequest)
    - [DeleteTaskResponse](#service-DeleteTaskResponse)
    - [Empty](#service-Empty)
    - [GetTaskRequest](#service-GetTaskRequest)
    - [GetTaskResponse](#service-GetTaskResponse)
    - [GetTasksResponse](#service-GetTasksResponse)
    - [Task](#service-Task)
    - [UpdateTaskRequest](#service-UpdateTaskRequest)
    - [UpdateTaskResponse](#service-UpdateTaskResponse)
  
    - [Service](#service-Service)
  
- [Scalar Value Types](#scalar-value-types)



<a name="protos_proto_Service_Service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## protos/proto/Service/Service.proto



<a name="service-CreateTaskRequest"></a>

### CreateTaskRequest
Запрос на создание задача


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |






<a name="service-CreateTaskResponse"></a>

### CreateTaskResponse
Ответ на запрос создания задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |
| status | [string](#string) |  | Статус задачи |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-DeleteTaskRequest"></a>

### DeleteTaskRequest
Запрос на удаление задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи которую хотим удалить |






<a name="service-DeleteTaskResponse"></a>

### DeleteTaskResponse
Ответ на запрос удаления задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Успешно ли удалена запись |






<a name="service-Empty"></a>

### Empty
Пустая структура, необходимая для запросов без входящих аргументов






<a name="service-GetTaskRequest"></a>

### GetTaskRequest
Запрос на извлечение задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи которую хотим получить |






<a name="service-GetTaskResponse"></a>

### GetTaskResponse
Ответ на запрос поиска задачи по ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |
| status | [string](#string) |  | Статус задачи |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-GetTasksResponse"></a>

### GetTasksResponse
Возвращает массив с списком всех задач


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [Task](#service-Task) | repeated | Массив задач |






<a name="service-Task"></a>

### Task
Структура задач


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |
| status | [string](#string) |  | Статус задачи |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |






<a name="service-UpdateTaskRequest"></a>

### UpdateTaskRequest
Запрос на обновление задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |
| status | [string](#string) |  | Статус задачи |






<a name="service-UpdateTaskResponse"></a>

### UpdateTaskResponse
Ответ на запрос обновлени задачи


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID задачи |
| title | [string](#string) |  | Название задачи |
| description | [string](#string) |  | Описание задачи |
| status | [string](#string) |  | Статус задачи |
| createdAt | [string](#string) |  | Дата создания сущности в базе данных |
| updatedAt | [string](#string) |  | Дата последнего обновления сущности в базе данных |





 

 

 


<a name="service-Service"></a>

### Service
Доступные API методы

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTask | [CreateTaskRequest](#service-CreateTaskRequest) | [CreateTaskResponse](#service-CreateTaskResponse) | Создание задачи |
| GetTasks | [Empty](#service-Empty) | [GetTasksResponse](#service-GetTasksResponse) | Список всех задач |
| GetTask | [GetTaskRequest](#service-GetTaskRequest) | [GetTaskResponse](#service-GetTaskResponse) | Получить определенную задачу по ID |
| UpdateTask | [UpdateTaskRequest](#service-UpdateTaskRequest) | [UpdateTaskResponse](#service-UpdateTaskResponse) | Обновление задачи |
| DeleteTask | [DeleteTaskRequest](#service-DeleteTaskRequest) | [DeleteTaskResponse](#service-DeleteTaskResponse) | Удаление задачи по ID |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

