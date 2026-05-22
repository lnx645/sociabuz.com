<script lang="ts">
  let isLoading = $state(true);
</script>

<!-- login page -->

<div class="max-w-sm mx-auto lg:p-0 p-2">
  <div class=" bg-surface-container shadow-surface-container p-6 mt-10 rounded-lg overflow-hidden">
    <div class="mb-4">
      <h2 class="text-2xl">Silahkan Login</h2>
      <p class="text-xs text-text-var">Silahkan login dan kelola penghasilan anda disini!</p>
    </div>
    <form action="">
      <div class="w-full mb-3">
        <label class="input-label" for="email">Email</label>
        <input
          placeholder="Ketikan email anda"
          class="input"
          type="text"
          id="email"
        />
      </div>
      <div class="w-full mb-3">
        <label class="input-label" for="password">Password</label>
        <input
          placeholder="Ketikan kata sandi anda"
          class="input"
          type="password"
          id="password"
        />
        <a class=" text-end block text-xs mt-1 mb-2  text-text-link" href="">
          Lupa Password?
        </a>
      </div>
      <div class="flex flex-col gap-2">
        <button disabled={isLoading} class="button">Login</button>
        <button disabled={isLoading} class="button auth__google-auth-btn"
          >Login Dengan Google</button
        >
      </div>
    </form>
  </div>
</div>
<!-- login page -->
