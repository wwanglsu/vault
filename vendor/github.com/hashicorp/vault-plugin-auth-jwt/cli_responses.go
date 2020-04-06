package jwtauth

import "fmt"

const successHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Vault Authentication Succeeded</title>
    <style>
      body {
        font-size: 14px;
        font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
          "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans",
          "Helvetica Neue", sans-serif;
      }
      hr {
        border-color: #fdfdfe;
        margin: 24px 0;
      }
      .container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 70vh;
      }
      #logo {
        display: block;
        fill: #6f7682;
        margin-bottom: 16px;
      }
      .message {
        display: flex;
        min-width: 40vw;
        background: #fafdfa;
        border: 1px solid #c6e9c9;
        margin-bottom: 12px;
        padding: 12px 16px 16px 12px;
        position: relative;
        border-radius: 2px;
        font-size: 14px;
      }
      .message-content {
        margin-left: 4px;
      }
      .message #checkbox {
        fill: #2eb039;
      }
      .message .message-title {
        color: #1e7125;
        font-size: 16px;
        font-weight: 700;
        line-height: 1.25;
      }
      .message .message-body {
        border: 0;
        margin-top: 4px;
      }
      .message p {
        font-size: 12px;
        margin: 0;
        padding: 0;
        color: #17421b;
      }
      a {
        display: block;
        margin: 8px 0;
        color: #1563ff;
        text-decoration: none;
        font-weight: 600;
      }
      a:hover {
        color: black;
      }
      a svg {
        fill: currentcolor;
      }
      .icon {
        align-items: center;
        display: inline-flex;
        justify-content: center;
        height: 21px;
        width: 21px;
        vertical-align: middle;
      }
      h1 {
        font-size: 17.5px;
        font-weight: 700;
        margin-bottom: 0;
      }
      h1 + p {
        margin: 8px 0 16px 0;
      }
    </style>
  </head>
  <body translate="no" >
    <div class="container">
      <div>
        <svg id="logo" fill="#000000" width="146" height="51" viewBox="0 0 146 51" xmlns="http://www.w3.org/2000/svg">
          <g id="vault-logo-v" fill-rule="nonzero">
            <path d=""></path>
          </g>
          <path id="vault-logo-name" d="M 42.414063 15 C 38.824219 15 36.574219 17.5 36.574219 17.5 C 35.378906 15.941406 33.730469 15.003906 30.941406 15.003906 C 27.996094 15.003906 26.042969 17.5 26.042969 17.5 C 24.847656 15.941406 22.6875 15 21 15 C 18.390625 15 16.320313 16.152344 15.054688 19.058594 L 10.820313 28.320313 L 6.03125 16.558594 C 5.425781 15.226563 3.933594 14.625 2.542969 15.246094 C 1.148438 15.871094 0.636719 17.425781 1.265625 18.757813 L 7.113281 31.945313 C 8.03125 33.949219 9.003906 35 10.820313 35 C 12.765625 35 13.609375 33.855469 14.53125 31.945313 C 14.53125 31.945313 18.511719 23.0625 19 22 C 19.488281 20.9375 20.300781 20 21.5 20 C 22.875 20 24 21.125 24 22.5 L 24 32.375 C 24 33.820313 25.085938 35 26.523438 35 C 27.957031 35 29 33.820313 29 32.375 L 29 22.5 C 29 21.125 30.125 20 31.5 20 C 32.875 20 34 21.125 34 22.5 L 34 32.5 C 34 33.875 35.125 35 36.5 35 C 37.875 35 39 33.875 39 32.5 L 39 22.5 C 39 21.125 40.125 20 41.5 20 C 42.875 20 44 21.125 44 22.5 L 44 32.5 C 44 33.875 45.125 35 46.5 35 C 47.875 35 49 33.875 49 32.5 L 49 21.355469 C 49 17.617188 46.011719 15 42.414063 15 Z"></path>
        </svg>
        <div class="message is-success">
          <svg id="checkbox" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 512 512">
            <path d="M256 32C132.3 32 32 132.3 32 256s100.3 224 224 224 224-100.3 224-224S379.7 32 256 32zm114.9 149.1L231.8 359.6c-1.1 1.1-2.9 3.5-5.1 3.5-2.3 0-3.8-1.6-5.1-2.9-1.3-1.3-78.9-75.9-78.9-75.9l-1.5-1.5c-.6-.9-1.1-2-1.1-3.2 0-1.2.5-2.3 1.1-3.2.4-.4.7-.7 1.1-1.2 7.7-8.1 23.3-24.5 24.3-25.5 1.3-1.3 2.4-3 4.8-3 2.5 0 4.1 2.1 5.3 3.3 1.2 1.2 45 43.3 45 43.3l111.3-143c1-.8 2.2-1.4 3.5-1.4 1.3 0 2.5.5 3.5 1.3l30.6 24.1c.8 1 1.3 2.2 1.3 3.5.1 1.3-.4 2.4-1 3.3z"></path>
        </svg>
          <div class="message-content">
            <div class="message-title">
              Signed in via your OIDC provider
            </div>
            <p class="message-body">
              You can now close this window and start using Vault.
            </p>
          </div>
        </div>
        <hr />
        <h1>Not sure how to get started?</h1>
        <p class="learn">
          Check out beginner and advanced guides on HashiCorp Vault at the HashiCorp Learn site or read more in the official documentation.
        </p>
        <a href="https://learn.hashicorp.com/vault" rel="noreferrer noopener">
         <span class="icon">
            <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
              <path d="M8.338 2.255a.79.79 0 0 0-.645 0L.657 5.378c-.363.162-.534.538-.534.875 0 .337.171.713.534.875l1.436.637c-.332.495-.638 1.18-.744 2.106a.887.887 0 0 0-.26 1.559c.02.081.03.215.013.392-.02.205-.074.43-.162.636-.186.431-.45.64-.741.64v.98c.651 0 1.108-.365 1.403-.797l.06.073c.32.372.826.763 1.455.763v-.98c-.215 0-.474-.145-.71-.42-.111-.13-.2-.27-.259-.393a1.014 1.014 0 0 1-.06-.155c-.01-.036-.013-.055-.013-.058h-.022a2.544 2.544 0 0 0 .031-.641.886.886 0 0 0-.006-1.51c.1-.868.398-1.477.699-1.891l.332.147-.023.746v2.228c0 .115.04.22.105.304.124.276.343.5.587.677.297.217.675.396 1.097.54.846.288 1.943.456 3.127.456 1.185 0 2.281-.168 3.128-.456.422-.144.8-.323 1.097-.54.244-.177.462-.401.586-.677a.488.488 0 0 0 .106-.304V8.218l2.455-1.09c.363-.162.534-.538.534-.875 0-.337-.17-.713-.534-.875L8.338 2.255zm-.34 2.955L3.64 7.38l4.375 1.942 6.912-3.069-6.912-3.07-6.912 3.07 1.665.74 4.901-2.44.328.657zM14.307 1H12.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L14.307 1zm-2.368 7.653v2.383a.436.436 0 0 0-.007.021c-.017.063-.084.178-.282.322-.193.14-.473.28-.836.404-.724.247-1.71.404-2.812.404-1.1 0-2.087-.157-2.811-.404a3.188 3.188 0 0 1-.836-.404c-.198-.144-.265-.26-.282-.322a.437.437 0 0 0-.007-.02V8.983l.01-.338 3.617 1.605a.791.791 0 0 0 .645 0l3.6-1.598z" fill-rule="evenodd"></path>
            </svg>
          </span>
          Get started with Vault
        </a>
        <a href="https://vaultproject.io/docs" rel="noreferrer noopener">
         <span class="icon">
          <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
    <path d="M13.307 1H11.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L13.307 1zM12 14V8a.5.5 0 1 1 1 0v6.5a.5.5 0 0 1-.5.5H.563a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .5-.5H8a.5.5 0 0 1 0 1H1v12h11zM4 6a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1H4zm0 2.5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1H4zM4 11a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1H4z"/>
  </svg> 
          </span>
          View the official Vault documentation
        </a>
      </div>
    </div>
  </body>
</html>
`

func errorHTML(summary, detail string) string {
	const html = `
<!DOCTYPE html>
<html lang="en" >

<head>

  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
<title>HashiCorp Vault</title>
      <style>
      body {
  font-size: 14px;
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
    "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans",
    "Helvetica Neue", sans-serif;
}
hr {
  border-color: #fdfdfe;
  margin: 24px 0;
}
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 70vh;
}
#logo {
  display: block;
  fill: #6f7682;
  margin-bottom: 16px;
}
.message {
  display: flex;
  min-width: 40vw;
  background: #fafdfa;
  border: 1px solid #c6e9c9;
  margin-bottom: 12px;
  padding: 12px 16px 16px 12px;
  position: relative;
  border-radius: 2px;
  font-size: 14px;
}
.message.is-danger {
  background: #fdfafb;
  border-color: #f9ecee;
}
.message-content {
  margin-left: 4px;
}
.message svg {
  fill: #2eb039;
}

.message.is-danger svg {
  fill: #c73445;
}
.message .message-title {
  color: #1e7125;
  font-size: 16px;
  font-weight: 700;
  line-height: 1.25;
}
.message.is-danger .message-title {
  color: #7f222c;
  
}
.message .message-body {
  border: 0;
  margin-top: 4px;
}
.message p {
  font-size: 12px;
  margin: 0;
  padding: 0;
  color: #17421b;
}
.message.is-danger p {
  color: #1f2124;
}
a {
  display: block;
  margin: 8px 0;
  color: #1563ff;
  text-decoration: none;
  font-weight: 600;
}
a:hover {
  color: black;
}
a svg {
  fill: currentcolor;
}
.icon {
  align-items: center;
  display: inline-flex;
  justify-content: center;
  height: 21px;
  width: 21px;
  vertical-align: middle;
}

h1 {
  font-size: 17.5px;
  font-weight: 700;
  margin-bottom: 0;
}

h1 + p {
  margin: 8px 0 16px 0;
}
    </style>
</head>
<body translate="no" >
  <div class="container">
    <div>
      <svg id="logo" fill="#000000" width="146" height="51" viewBox="0 0 146 51" xmlns="http://www.w3.org/2000/svg">
  <g id="vault-logo-v" fill-rule="nonzero">
    <path d=""></path>
  </g>
  <path id="vault-logo-name" d="M 42.414063 15 C 38.824219 15 36.574219 17.5 36.574219 17.5 C 35.378906 15.941406 33.730469 15.003906 30.941406 15.003906 C 27.996094 15.003906 26.042969 17.5 26.042969 17.5 C 24.847656 15.941406 22.6875 15 21 15 C 18.390625 15 16.320313 16.152344 15.054688 19.058594 L 10.820313 28.320313 L 6.03125 16.558594 C 5.425781 15.226563 3.933594 14.625 2.542969 15.246094 C 1.148438 15.871094 0.636719 17.425781 1.265625 18.757813 L 7.113281 31.945313 C 8.03125 33.949219 9.003906 35 10.820313 35 C 12.765625 35 13.609375 33.855469 14.53125 31.945313 C 14.53125 31.945313 18.511719 23.0625 19 22 C 19.488281 20.9375 20.300781 20 21.5 20 C 22.875 20 24 21.125 24 22.5 L 24 32.375 C 24 33.820313 25.085938 35 26.523438 35 C 27.957031 35 29 33.820313 29 32.375 L 29 22.5 C 29 21.125 30.125 20 31.5 20 C 32.875 20 34 21.125 34 22.5 L 34 32.5 C 34 33.875 35.125 35 36.5 35 C 37.875 35 39 33.875 39 32.5 L 39 22.5 C 39 21.125 40.125 20 41.5 20 C 42.875 20 44 21.125 44 22.5 L 44 32.5 C 44 33.875 45.125 35 46.5 35 C 47.875 35 49 33.875 49 32.5 L 49 21.355469 C 49 17.617188 46.011719 15 42.414063 15 Z"></path>
</svg>
      <div class="message is-danger">
       <svg width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
  <path d="M19 3c1.1 0 2 .9 2 2v14c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2V5c0-1.1.9-2 2-2h14zm-2 12.59L13.41 12 17 8.41 15.59 7 12 10.59 8.41 7 7 8.41 10.59 12 7 15.59 8.41 17 12 13.41 15.59 17 17 15.59z"></path>
</svg> 
        <div class="message-content">
          <div class="message-title">
            %s
          </div>
          <p class="message-body">
            %s
          </p>
        </div>
      </div>
      <hr />

      <h1>Not sure how to get started?</h1>
      <p class="learn">
        Check out beginner and advanced guides on HashiCorp Vault at the HashiCorp Learn site or read more in the official documentation.
      </p>
      <a href="https://learn.hashicorp.com/vault" rel="noreferrer noopener">
       <span class="icon">
      <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
  <path d="M8.338 2.255a.79.79 0 0 0-.645 0L.657 5.378c-.363.162-.534.538-.534.875 0 .337.171.713.534.875l1.436.637c-.332.495-.638 1.18-.744 2.106a.887.887 0 0 0-.26 1.559c.02.081.03.215.013.392-.02.205-.074.43-.162.636-.186.431-.45.64-.741.64v.98c.651 0 1.108-.365 1.403-.797l.06.073c.32.372.826.763 1.455.763v-.98c-.215 0-.474-.145-.71-.42-.111-.13-.2-.27-.259-.393a1.014 1.014 0 0 1-.06-.155c-.01-.036-.013-.055-.013-.058h-.022a2.544 2.544 0 0 0 .031-.641.886.886 0 0 0-.006-1.51c.1-.868.398-1.477.699-1.891l.332.147-.023.746v2.228c0 .115.04.22.105.304.124.276.343.5.587.677.297.217.675.396 1.097.54.846.288 1.943.456 3.127.456 1.185 0 2.281-.168 3.128-.456.422-.144.8-.323 1.097-.54.244-.177.462-.401.586-.677a.488.488 0 0 0 .106-.304V8.218l2.455-1.09c.363-.162.534-.538.534-.875 0-.337-.17-.713-.534-.875L8.338 2.255zm-.34 2.955L3.64 7.38l4.375 1.942 6.912-3.069-6.912-3.07-6.912 3.07 1.665.74 4.901-2.44.328.657zM14.307 1H12.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L14.307 1zm-2.368 7.653v2.383a.436.436 0 0 0-.007.021c-.017.063-.084.178-.282.322-.193.14-.473.28-.836.404-.724.247-1.71.404-2.812.404-1.1 0-2.087-.157-2.811-.404a3.188 3.188 0 0 1-.836-.404c-.198-.144-.265-.26-.282-.322a.437.437 0 0 0-.007-.02V8.983l.01-.338 3.617 1.605a.791.791 0 0 0 .645 0l3.6-1.598z" fill-rule="evenodd"></path>
</svg>
        </span>
  Get started with Vault
      </a>

      
      <a href="https://vaultproject.io/docs" rel="noreferrer noopener">
       <span class="icon">
        <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
  <path d="M13.307 1H11.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L13.307 1zM12 14V8a.5.5 0 1 1 1 0v6.5a.5.5 0 0 1-.5.5H.563a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .5-.5H8a.5.5 0 0 1 0 1H1v12h11zM4 6a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1H4zm0 2.5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1H4zM4 11a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1H4z"/>
</svg> 
        </span>
        View the official Vault documentation
      </a>
    </div>
  </div>
</body>

</html>
`
	return fmt.Sprintf(html, summary, detail)
}

func formpostHTML(path, code, state string) string {
	const html = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Complete sign-in process</title>
    <style>
      body {
        font-size: 14px;
        font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI",
          "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans",
          "Helvetica Neue", sans-serif;
      }
      hr {
        border-color: #fdfdfe;
        margin: 24px 0;
      }
      .container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 70vh;
      }
      #logo {
        display: block;
        fill: #6f7682;
        margin-bottom: 16px;
      }
      .message {
        display: flex;
        min-width: 40vw;
        background: #fafdfa;
        border: 1px solid #c6e9c9;
        margin-bottom: 12px;
        padding: 12px 16px 16px 12px;
        position: relative;
        border-radius: 2px;
        font-size: 14px;
      }
      .message-content {
        margin-left: 4px;
      }
      .message #checkbox {
        fill: #2eb039;
      }
      .message .message-title {
        color: #1e7125;
        font-size: 16px;
        font-weight: 700;
        line-height: 1.25;
      }
      .message .message-body {
        border: 0;
        margin-top: 4px;
      }
      .message p {
        font-size: 12px;
        margin: 0;
        padding: 0;
        color: #17421b;
      }
      a {
        display: block;
        margin: 8px 0;
        color: #1563ff;
        text-decoration: none;
        font-weight: 600;
      }
      a:hover {
        color: black;
      }
      a svg {
        fill: currentcolor;
      }
      .icon {
        align-items: center;
        display: inline-flex;
        justify-content: center;
        height: 21px;
        width: 21px;
        vertical-align: middle;
      }
      h1 {
        font-size: 17.5px;
        font-weight: 700;
        margin-bottom: 0;
      }
      h1 + p {
        margin: 8px 0 16px 0;
      }
    </style>
  </head>
  <body translate="no" >
    <div class="container">
      <div>
        <svg id="logo" fill="#000000" width="146" height="51" viewBox="0 0 146 51" xmlns="http://www.w3.org/2000/svg">
          <g id="vault-logo-v" fill-rule="nonzero">
            <path d=""></path>
          </g>
          <path id="vault-logo-name" d="M 42.414063 15 C 38.824219 15 36.574219 17.5 36.574219 17.5 C 35.378906 15.941406 33.730469 15.003906 30.941406 15.003906 C 27.996094 15.003906 26.042969 17.5 26.042969 17.5 C 24.847656 15.941406 22.6875 15 21 15 C 18.390625 15 16.320313 16.152344 15.054688 19.058594 L 10.820313 28.320313 L 6.03125 16.558594 C 5.425781 15.226563 3.933594 14.625 2.542969 15.246094 C 1.148438 15.871094 0.636719 17.425781 1.265625 18.757813 L 7.113281 31.945313 C 8.03125 33.949219 9.003906 35 10.820313 35 C 12.765625 35 13.609375 33.855469 14.53125 31.945313 C 14.53125 31.945313 18.511719 23.0625 19 22 C 19.488281 20.9375 20.300781 20 21.5 20 C 22.875 20 24 21.125 24 22.5 L 24 32.375 C 24 33.820313 25.085938 35 26.523438 35 C 27.957031 35 29 33.820313 29 32.375 L 29 22.5 C 29 21.125 30.125 20 31.5 20 C 32.875 20 34 21.125 34 22.5 L 34 32.5 C 34 33.875 35.125 35 36.5 35 C 37.875 35 39 33.875 39 32.5 L 39 22.5 C 39 21.125 40.125 20 41.5 20 C 42.875 20 44 21.125 44 22.5 L 44 32.5 C 44 33.875 45.125 35 46.5 35 C 47.875 35 49 33.875 49 32.5 L 49 21.355469 C 49 17.617188 46.011719 15 42.414063 15 Z"></path>
        </svg>
        <div class="message is-success">
          <svg id="checkbox" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 512 512">
            <path d="M256 32C132.3 32 32 132.3 32 256s100.3 224 224 224 224-100.3 224-224S379.7 32 256 32zm114.9 149.1L231.8 359.6c-1.1 1.1-2.9 3.5-5.1 3.5-2.3 0-3.8-1.6-5.1-2.9-1.3-1.3-78.9-75.9-78.9-75.9l-1.5-1.5c-.6-.9-1.1-2-1.1-3.2 0-1.2.5-2.3 1.1-3.2.4-.4.7-.7 1.1-1.2 7.7-8.1 23.3-24.5 24.3-25.5 1.3-1.3 2.4-3 4.8-3 2.5 0 4.1 2.1 5.3 3.3 1.2 1.2 45 43.3 45 43.3l111.3-143c1-.8 2.2-1.4 3.5-1.4 1.3 0 2.5.5 3.5 1.3l30.6 24.1c.8 1 1.3 2.2 1.3 3.5.1 1.3-.4 2.4-1 3.3z"></path>
        </svg>
          <div class="message-content">
            <div class="message-title">
              Completing the sign-in process...
            </div>
          </div>
        </div>
        <hr />
        <h1>Not sure how to get started?</h1>
        <p class="learn">
          Check out beginner and advanced guides on HashiCorp Vault at the HashiCorp Learn site or read more in the official documentation.
        </p>
        <a href="https://learn.hashicorp.com/vault" rel="noreferrer noopener">
         <span class="icon">
            <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
              <path d="M8.338 2.255a.79.79 0 0 0-.645 0L.657 5.378c-.363.162-.534.538-.534.875 0 .337.171.713.534.875l1.436.637c-.332.495-.638 1.18-.744 2.106a.887.887 0 0 0-.26 1.559c.02.081.03.215.013.392-.02.205-.074.43-.162.636-.186.431-.45.64-.741.64v.98c.651 0 1.108-.365 1.403-.797l.06.073c.32.372.826.763 1.455.763v-.98c-.215 0-.474-.145-.71-.42-.111-.13-.2-.27-.259-.393a1.014 1.014 0 0 1-.06-.155c-.01-.036-.013-.055-.013-.058h-.022a2.544 2.544 0 0 0 .031-.641.886.886 0 0 0-.006-1.51c.1-.868.398-1.477.699-1.891l.332.147-.023.746v2.228c0 .115.04.22.105.304.124.276.343.5.587.677.297.217.675.396 1.097.54.846.288 1.943.456 3.127.456 1.185 0 2.281-.168 3.128-.456.422-.144.8-.323 1.097-.54.244-.177.462-.401.586-.677a.488.488 0 0 0 .106-.304V8.218l2.455-1.09c.363-.162.534-.538.534-.875 0-.337-.17-.713-.534-.875L8.338 2.255zm-.34 2.955L3.64 7.38l4.375 1.942 6.912-3.069-6.912-3.07-6.912 3.07 1.665.74 4.901-2.44.328.657zM14.307 1H12.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L14.307 1zm-2.368 7.653v2.383a.436.436 0 0 0-.007.021c-.017.063-.084.178-.282.322-.193.14-.473.28-.836.404-.724.247-1.71.404-2.812.404-1.1 0-2.087-.157-2.811-.404a3.188 3.188 0 0 1-.836-.404c-.198-.144-.265-.26-.282-.322a.437.437 0 0 0-.007-.02V8.983l.01-.338 3.617 1.605a.791.791 0 0 0 .645 0l3.6-1.598z" fill-rule="evenodd"></path>
            </svg>
          </span>
          Get started with Vault
        </a>
        <a href="https://vaultproject.io/docs" rel="noreferrer noopener">
         <span class="icon">
          <svg width="16" height="16" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
    <path d="M13.307 1H11.5a.5.5 0 1 1 0-1h3a.499.499 0 0 1 .5.65V3.5a.5.5 0 1 1-1 0V1.72l-1.793 1.774a.5.5 0 0 1-.713-.701L13.307 1zM12 14V8a.5.5 0 1 1 1 0v6.5a.5.5 0 0 1-.5.5H.563a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 .5-.5H8a.5.5 0 0 1 0 1H1v12h11zM4 6a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1H4zm0 2.5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1H4zM4 11a.5.5 0 1 1 0-1h5a.5.5 0 1 1 0 1H4z"/>
  </svg> 
          </span>
          View the official Vault documentation
        </a>
      </div>
    </div>
	<script>
		window.localStorage.setItem("oidcState", JSON.stringify({"path":"%s", "code":"%s", "state":"%s"}));
	</script>
  </body>
</html>
`
	return fmt.Sprintf(html, path, code, state)
}
