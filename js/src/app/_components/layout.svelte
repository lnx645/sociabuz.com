<script lang="ts">
  const { children } = $props();
</script>

<nav class="w-full h-14 bg-navbar shadow-1">
  <div class="container mx-auto px-24 h-14 flex items-center">
    <a href="/" class="font-bold text-lg">Sobuzz.id</a>
    <div class="ml-auto">
      <a href="/login" class="text-sm text-text-1">Login</a>
      <a href="/register" class="text-sm text-text-1 ml-4">Register</a>
    </div>
  </div>
</nav>
<main>
  {@render children?.()}
</main>
