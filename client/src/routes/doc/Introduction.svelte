<script></script>

<div class="Introduction-container">
  <div class="title-container">
    <h1>Introduction</h1>
    <p>
      Lapis OAuth is an authentication system utilizing the OAuth 2.0 standards
      to deliver a secure client-developer relationship. This OAuth application
      is written in Go and utilizes Redis as a cache system and PostgreSQL as a
      database.
      <br /><br />
      The application provides an access and refresh token on successful authentication
      by creating a time-bound cookie.
    </p>
  </div>
  <div class="Authentication-container sub-container">
    <h3>Authentication</h3>
    <p>
      For the client, the authentication is as simple as it gets. Registration
      requires some user data, and login is done using a password method.
      <br /><br />
      For the developer, they can request user data using a valid access token. If
      an access token expires, the developers can use the refresh token to get a
      new access token. Every time a refresh token is used, it becomes invalid, and
      a new refresh token is generated in its place.
    </p>
  </div>
  <div class="how-container sub-container">
    <h3>How To Use?</h3>
    <p>
      To use the OAuth application, you can simply redirect to the “/authorize”
      endpoint. It will create a new OAuth session in the browser in a separate
      window where the user can log in or sign up.
      <br /><br />
      Once await authenticated, a cookie is generated with the access token, refresh
      token, and user ID. This data can be used to access user data while valid.
      <br /><br />
      To view the full API endpoint details, please refer to the routes section.
    </p>
  </div>
</div>
