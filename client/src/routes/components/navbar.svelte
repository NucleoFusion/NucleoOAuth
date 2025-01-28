<script>
  import { goto } from "$app/navigation";
</script>

<div class="navbar">
  <div>
    <button
      onclick={() => {
        goto("/");
      }}>Home</button
    >
  </div>
  <div>
    <button
      onclick={() => {
        goto("/doc");
      }}>Doc</button
    >
  </div>
</div>

<style>
  .navbar {
    position: fixed;

    top: 3vh;
    left: 32vw;

    width: 30vw;
    height: 10vh;

    padding: 1vh 3vw;

    border-radius: 100px;

    background-color: #b4befe;

    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr;
  }

  .navbar > div {
    width: 100%;
    height: 100%;

    display: grid;
    justify-content: center;
    align-items: center;
    text-align: center;
  }

  .navbar button {
    width: 7.5vw;
    height: 5vh;

    border: none;
    border-radius: 15px;

    background-color: #b4befe;
    color: #1e1e2e;

    border: solid #1e1e2e 1px;

    font-size: 1.25em;
  }

  .navbar button:hover {
    background-color: #1e1e2e;
    color: #b4befe;

    cursor: pointer;

    border: solid #1e1e2e 1px;
  }

  @media (max-width: 1000px) {
    .navbar button {
      width: 15vw;
    }

    .navbar {
      width: 45vw;
      left: 24.5vw;
    }
  }

  @media (max-width: 700px) {
    .navbar button {
      width: 15vw;
    }

    .navbar {
      width: 55vw;
      left: 19.5vw;
    }
  }

  @media (max-width: 500px) {
    .navbar button {
      width: 25vw;
    }

    .navbar {
      width: 75vw;
      left: 9.5vw;
    }
  }
</style>
