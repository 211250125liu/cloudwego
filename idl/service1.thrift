namespace go demo

struct Request {
    1: string message(api.body = 'message')
}

struct Response {
    1: string message(api.body = "message")
}

service service_1{
    Response getMessage(1: Request req)(api.get = '/getMessage')
}