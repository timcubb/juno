package deploy

type App struct {
    Name string
    Repo string
    Image string
}

type Deploy struct {
    App App
    Node Node
    Instances map[string]AppInstance
}