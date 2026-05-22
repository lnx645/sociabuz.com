<script lang="ts">
  let isLoading = $state(true);
</script>

<!-- login page -->

<div class="auth">
  <div class="auth__login">
    <div class="auth__header">
      <h2>Silahkan Login</h2>
    </div>
    <form action="">
      <div class="auth__group-field">
        <label for="email">Email</label>
        <input
          placeholder="Ketikan email anda"
          class="input"
          type="text"
          id="email"
        />
      </div>
      <div class="auth__group-field">
        <label for="katasandi">Kata Sandi</label>
        <input
          placeholder="Ketikan email anda"
          class="input"
          type="password"
          id="email"
        />
      </div>
      <div class="">
        <button disabled={isLoading} class="button">Login</button>
      </div>
    </form>
  </div>
</div>
<!-- login page -->
