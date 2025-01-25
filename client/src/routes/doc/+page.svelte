<script>
  import jq from "jquery";

  import Introduction from "./Introduction.svelte";
  import Routes from "./Routes.svelte";
  import About from "./About.svelte";

  let current = $state("Introduction");

  $effect(() => {
    if (
      current != "Introduction" &&
      current != "About" &&
      current != "Routes"
    ) {
      current = "Introduction";
    }

    jq(`#${current}>button`).attr("disabled", "disabled");
  });
</script>

<div class="doc-container jetbrains-mono">
  <div class="navigator">
    <div>Lapis OAuth</div>
    <hr />
    <div id="Introduction">
      <button
        disabled
        onclick={() => {
          current = "Introduction";

          jq(".navigator>div>button:disabled").removeAttr("disabled");
        }}>Introduction</button
      >
    </div>
    <div id="Routes">
      <button
        onclick={() => {
          current = "Routes";

          jq(".navigator>div>button:disabled").removeAttr("disabled");
        }}>Routes</button
      >
    </div>
    <div id="About">
      <button
        onclick={() => {
          current = "About";

          jq(".navigator>div>button:disabled").removeAttr("disabled");
        }}>About</button
      >
    </div>
  </div>
  {#key current}
    {#if current == "Introduction"}
      <Introduction />
    {:else if current == "Routes"}
      <Routes />
    {:else if current == "About"}
      <About />
    {/if}
  {/key}
</div>

<style>
  @import "./doc.css";
</style>
