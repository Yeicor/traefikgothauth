package traefikgothauth

import "html/template"

//go:generate /usr/bin/env bash -c "cd assets/ && npx parcel build login-choose-provider.html && sed -i -E 's/<!--|-->//g' dist/login-choose-provider.html && ( awk 'NR < 8 { print }' ../login.go && printf 'var loginChooseProviderHtml = template.Must(template.New(`loginChooseProviderTemplate`).Parse(`' && cat dist/login-choose-provider.html && printf '`))' ) >_tmp.go && mv _tmp.go ../login.go"

// DO-NOT-EDIT ANYTHING IN THIS FILE (EVERYTHING BELOW THIS LINE WILL BE DELETED)
var loginChooseProviderHtml = template.Must(template.New(`loginChooseProviderTemplate`).Parse(`<!DOCTYPE html><html lang="en"><head><title>Login</title><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"></head><body> <div class="login-box"> <div class="login-box-title">Login</div> <div class="login-box-buttons"> {{range $key,$value:=.}} <a class="btn" href="/__goth/{{$value.Name}}/login/" > <img src="{{$value.Icon}}"  alt="icon"> {{$value.DisplayName}} </a> {{end}} </div> </div> <style>html,body{margin:0;padding:0;font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Helvetica Neue,Arial,sans-serif;font-size:16px}.login-box{border-radius:20px;min-width:300px;max-width:-moz-fit-content;max-width:fit-content;margin-top:5vh;margin-left:auto;margin-right:auto;padding:20px;box-shadow:0 0 10px #0000001a}.login-box-title{text-align:center;margin-bottom:20px;font-size:24px;font-weight:700}.login-box-buttons{flex-flow:wrap;justify-content:center;align-items:center;display:flex}.btn{color:#000;background-color:#f0f0f0;border-radius:5px;margin:10px;padding:10px 20px;text-decoration:none;transition:background-color .3s;display:inline-block}.btn:hover{background-color:#e0e0e0}.btn img{vertical-align:text-top;width:20px;margin-right:5px}@media (prefers-color-scheme:dark){body{color:#f0f0f0;background-color:#121212}.login-box{color:#f0f0f0;background-color:#1e1e1e}.btn{color:#f0f0f0;background-color:#333}.btn:hover{background-color:#444}}@media (width>=1200px){.login-box{max-width:1200px}}</style> </body></html>`))