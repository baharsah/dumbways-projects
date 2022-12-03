module.exports = {
  apps : [{
    name   : "golang",
    script : "go run main.go",
    watch : ["./main.go" , "controller" , "collector" , "router"]
  },
  {
    name   : "scss",
    script : "npm run scss"
  }]
}
