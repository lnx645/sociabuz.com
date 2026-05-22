<script lang="ts">
  const { children } = $props();
</script>

<nav class="navbar">
  <div class="navbar__wrapper">
    <a href="/" class="navbar__brand">Sobuzz</a>
    <div class="navbar__right">
      <a href="/login" class="navbar__link">Login</a>
      <a href="/register" class="navbar__link">Register</a>
    </div>
  </div>
</nav>
<main>
  {@render children?.()}
  
</main>
