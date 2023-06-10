<script>
  let username = '';
  let password = '';

  function handleSubmit() {
    const authHeader = 'Basic ' + btoa(`${username}:${password}`);
    // Basicauthでデータを送信する処理をここに追加する
    console.log('Auth Header:', authHeader);
  }
</script>

<div>
  <label for="username">ID:</label>
  <input type="text" id="username" bind:value={username} />

  <label for="password">パスワード:</label>
  <input type="password" id="password" bind:value={password} />

  <button on:click={handleSubmit}>送信</button>
</div>

<style>
  /* スタイリングを適用する場合はここに記述 */
</style>
