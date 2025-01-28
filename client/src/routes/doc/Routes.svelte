<script></script>

<div class="Routes-container">
  <h1>Routes</h1>
  <div class="Routes-introduction">
    <p>
      Here you can find details about the client facing routes in our OAuth
      application.
    </p>
  </div>
  <div class="route">
    <h2>
      <span class="method"> GET </span>
      /authorize
    </h2>
    <p class="route-descr">
      This route is what developers can use to open up a new OAuth session in
      the user's browser.
    </p>
    <div class="response">
      <h3>Response</h3>
      <p>No response is generated. Redirects to the login page.</p>
    </div>
    <div class="params">
      <h3>Parameters</h3>
      <div class="path-params">
        <h4>Path</h4>
        <p>
          <span class="name-params"> No Path Params </span>
        </p>
      </div>
      <div class="query-params">
        <h4>Query</h4>
        <p>
          <span class="name-params"> No Query Params </span>
        </p>
      </div>
      <div class="body-params">
        <h4>Body</h4>
        <p>
          <span class="name-params"> No Body Params </span>
        </p>
      </div>
    </div>
  </div>
  <div class="route">
    <h2>
      <span class="method"> POST </span>
      /refresh
    </h2>
    <p class="route-descr">
      This Route is used to get a new Access Token by using a Refresh Token.
    </p>
    <div class="response">
      <h3>Response</h3>
      <p>
        user_id int
        <br />
        access_token string
        <br />
        refresh_token string
        <br />
        type string
        <br />
        expires_in int
      </p>
    </div>
    <div class="params">
      <h3>Parameters</h3>
      <div class="path-params">
        <h4>Path</h4>
        <p>
          <span class="name-params"> No Path Params </span>
        </p>
      </div>
      <div class="query-params">
        <h4>Query</h4>
        <p>
          <span class="name-params"> No Query Params </span>
        </p>
      </div>
      <div class="body-params">
        <h4>Body</h4>
        <p>
          <span class="name-params"> refresh_token string </span>
        </p>
      </div>
    </div>
  </div>
  <div class="route">
    <h2>
      <span class="method"> POST </span>
      /api/access
    </h2>
    <p class="route-descr">
      This route can be used to gain access to a specific user's details by
      using an Access Token.
      <br />
      The request must be made to the resource server url.
    </p>
    <div class="response">
      <h3>Response</h3>
      <p>
        user_id string
        <br />
        address string
        <br />
        phone int
        <br />
        country string
        <br />
        state string
        <br />
        name string
        <br />
        email string
      </p>
    </div>
    <div class="params">
      <h3>Parameters</h3>
      <div class="path-params">
        <h4>Path</h4>
        <p>
          <span class="name-params"> No Path Params </span>
        </p>
      </div>
      <div class="query-params">
        <h4>Query</h4>
        <p>
          <span class="name-params">
            access_token string
            <br />
            user_id int
          </span>
        </p>
      </div>
      <div class="body-params">
        <h4>Body</h4>
        <p>
          <span class="name-params"> No Body Params </span>
        </p>
      </div>
    </div>
  </div>
</div>
