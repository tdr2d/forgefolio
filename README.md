## Vue - CLI
- Install vue cli globally:  `sudo npm install -g @vue/cli @vue/cli-service-global`
- Create new project `vue create my-project`
- Add typscript support to the project `vue add typescript`

## Go Quicktemplate : https://github.com/valyala/quicktemplate
- dependencies 
    ```
    go get -u github.com/valyala/quicktemplate
    go get -u github.com/valyala/quicktemplate/qtc
    ```
- build templates `make gen`
- on vscode: settings > file:exclude > add `**/*.qtpl.go`