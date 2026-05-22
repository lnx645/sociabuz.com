<script lang="ts">
  const { children } = $props();
</script>

<nav class="w-full h-14 bg-primary shadow-header text-white">
  <div class="flex container px-3 lg:px-14 mx-auto h-14 items-center">
    <a href="/" class="navbar__brand">Sobuzz</a>
    <div class="ml-auto">
      <a href="/login" class="text-sm">Login</a>
      <a href="/register" class="text-sm">Register</a>
    </div>
  </div>
</nav>
<main>
  {@render children?.()}
</main>
