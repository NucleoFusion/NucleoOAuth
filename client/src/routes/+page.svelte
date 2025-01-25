<script>
  import { goto } from "$app/navigation";
  import Navbar from "./components/navbar.svelte";
</script>

<Navbar />
<div class="home-container">
  <div class="home-heading">
    <h1 class="jetbrains-mono">Lapis OAuth</h1>
  </div>
  <div class="home-description jetbrains-mono">
    <p>An OAuth application written in Go, Postgres and Redis.</p>
  </div>
  <div class="home-links">
    <div>
      <button
        class="jetbrains-mono"
        onclick={() => {
          goto("/doc");
        }}>Docs</button
      >
    </div>
    <div>
      <button
        class="jetbrains-mono"
        onclick={() => {
          window.location = "https://github.com/NucleoFusion/NucleoOAuth";
        }}>GitHub</button
      >
    </div>
    <div>
      <button
        class="jetbrains-mono"
        onclick={() => {
          goto("/");
        }}>About</button
      >
    </div>
  </div>
</div>

<style>
  .home-container {
    width: 60vw;
    height: 76vh;

    margin-top: 15vh;
    margin-left: 18vw;

    display: grid;

    grid-template-columns: 1fr;
    grid-template-rows: 1fr 0.25fr 1fr;

    padding: 2vh 1vw;
  }

  .home-container > div {
    width: 100%;
    height: 100%;

    display: grid;

    justify-content: center;
    align-items: center;

    color: #b4befe;
  }

  .home-heading > h1 {
    margin: 0;

    font-size: 12vh;
  }

  .home-description > p {
    width: 40vw;

    text-align: center;
  }

  .home-links {
    display: grid;

    grid-template-rows: 1fr;
    grid-template-columns: 1fr 1fr 1fr;
  }

  .home-links button {
    width: 10vw;
    height: 4vh;

    font-size: 2svh;

    background-color: #b4befe;
    border: none;
    border-radius: 10px;
  }

  .home-links button:hover {
    background-color: rgb(180, 190, 254, 0.8);
    cursor: pointer;
  }

  .home-links > div {
    display: grid;

    justify-content: center;
    align-items: center;
    text-align: center;
  }
</style>
