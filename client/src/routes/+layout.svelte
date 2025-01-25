<script>
  let { children } = $props();
</script>

{@render children()}

<style>
  :global(body) {
    background-color: #1e1e2e;
  }

  :global(button) {
    transition: all 0.3s ease-out;
  }

  :global(.jetbrains-mono) {
    font-family: "JetBrains Mono", serif;
    font-optical-sizing: auto;
    font-weight: 400;
    font-style: normal;
  }
</style>
