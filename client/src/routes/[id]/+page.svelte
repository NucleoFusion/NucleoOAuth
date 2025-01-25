<script>
  import { page } from "$app/stores";
  import axios from "axios";

  import jq from "jquery";

  import { PUBLIC_API_URL } from "$env/static/public";

  let id = $page.params.id;
  let error = $state("");

  $effect(() => {
    if (error != "") {
      jq(".popup").focus();
    }
  });

  async function formSubmit(e) {
    e.preventDefault();

    let url = `${PUBLIC_API_URL}/`;

    if (isLogin) {
      url += `login/${id}`;
    } else {
      url += `register/${id}`;
    }

    const headers = {
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
    };

    const data = {
      email: jq("input[name='email']").val(),
      name: jq("input[name='name']").val(),
      pass: jq("input[name='pass']").val(),
    };

    const result = await axios.post(url, data, headers);

    if (result.data.error) {
      error = result.data.error;
    }
  }

  let isLogin = $state(true);

  function popdown(e) {
    console.log(e);
  }

  let onclick = () => {
    isLogin = !isLogin;
  };
</script>

<div class="main-container">
  <div class="title jetbrains-mono">
    <h1>Lapis OAuth</h1>
  </div>
  {#key isLogin}
    <div class="form-container">
      <form onsubmit={formSubmit}>
        {#if !isLogin}
          <div class="input-container">
            <label class="jetbrains-mono" for="name">Name: </label>
            <input class="jetbrains-mono" type="text" id="name" name="name" />
          </div>
        {/if}
        <div class="input-container">
          <label class="jetbrains-mono" for="email">Email: </label>
          <input class="jetbrains-mono" type="email" id="email" name="email" />
        </div>
        <div class="input-container">
          <label class="jetbrains-mono" for="pass">Password: </label>
          <input class="jetbrains-mono" type="password" id="pass" name="pass" />
        </div>
        <div class="submit-container">
          <button class="submit-button jetbrains-mono" type="submit">
            {#if isLogin}
              Login
            {:else}
              Sign Up
            {/if}
          </button>
        </div>
      </form>
    </div>
    <div class="button-container">
      <button class="jetbrains-mono" {onclick}>
        {#if isLogin}
          Sign Up
        {:else}
          Login
        {/if}
      </button>
    </div>
  {/key}
  {#key error}
    {#if error != ""}
      <div
        class="popup"
        tabindex="-1"
        onkeypress={(e) => {
          e.keyCode == 13 ? (error = "") : error;
        }}
      >
        <h3>{error}</h3>
      </div>
    {/if}
  {/key}
</div>

<style>
  :global(body) {
    padding: 0;
  }

  .main-container {
    width: 90vw;
    height: 90vh;

    margin: 5vh 5vw;

    display: grid;
    grid-template-rows: 1fr 4fr;
    grid-template-columns: 1fr 1fr;
  }

  .main-container > div {
    display: grid;
    justify-content: center;
    align-items: center;
  }

  .title {
    grid-column: span 2;
    font-size: 2em;
  }

  .button-container {
    padding: 15vh 3vw;
  }

  .button-container button {
    width: 15vw;
    height: 5vh;

    font-size: 1em;
  }

  .form-container form {
    width: 45vw;
    height: 72vh;

    display: grid;
    grid-template-rows: 1fr 1fr 1fr 1fr;
    grid-template-columns: 1fr;

    justify-content: center;
    align-items: center;
  }

  .input-container {
    padding: 3vh 3vw;

    display: grid;
    grid-template-columns: 1fr 3fr;
    grid-template-rows: 1fr;

    justify-content: center;
    align-items: center;
    text-align: center;
  }

  .input-container input {
    height: 3vh;
  }

  .submit-container {
    width: 100%;
    height: 100%;

    display: grid;
    justify-content: center;
    align-items: center;
    text-align: center;
  }

  .submit-button {
    width: 15vw;
    height: 5vh;

    font-size: 1em;
  }

  .popup {
    width: 50vw;
    height: 50vh;

    left: 25vw;
    top: 25vh;

    background-color: gray;

    position: fixed;

    z-index: 11;
  }
</style>
