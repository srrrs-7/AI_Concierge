<script>
    let name = '';
    let email = '';

    function handleSubmit() {
      alert(`名前: ${name}\nメールアドレス: ${email}`);
    }
  </script>

<head>
    <meta charset="UTF-8">
</head>


  <main>
    <h1>form メールアドレス</h1>

    <form on:submit|preventDefault={handleSubmit}>
      <label>
        名前:
        <input type="text" bind:value={name} />
      </label>

      <label>
        メールアドレス:
        <input type="email" bind:value={email} />
      </label>

      <button type="submit">send</button>
    </form>
  </main>

  <style>
    main {
      text-align: center;
      padding: 1em;
      max-width: 240px;
      margin: 0 auto;
    }

    input {
      display: block;
      margin-bottom: 0.5em;
    }

    button {
      padding: 0.5em 1em;
    }
  </style>
