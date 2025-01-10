<script>
  import { page } from "$app/stores";
  let id = $page.params.id;

  async function formSubmit(e) {
    e.preventDefault();
  }

  let isLogin = $state(true);

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
      <form action={formSubmit}>
        {#if !isLogin}
          <div class="input-container">
            <label class="jetbrains-mono" for="name">Name: </label>
            <input class="jetbrains-mono" type="text" id="name" />
          </div>
        {/if}
        <div class="input-container">
          <label class="jetbrains-mono" for="email">Email: </label>
          <input class="jetbrains-mono" type="email" id="email" />
        </div>
        <div class="input-container">
          <label class="jetbrains-mono" for="pass">Password: </label>
          <input class="jetbrains-mono" type="password" id="pass" />
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
</div>

<style>
  :global(body) {
    padding: 0;
  }

  :global(.jetbrains-mono) {
    font-family: "JetBrains Mono", serif;
    font-optical-sizing: auto;
    font-weight: 400;
    font-style: normal;
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
</style>
