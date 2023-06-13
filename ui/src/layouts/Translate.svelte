<script>
  import { t, locale, locales } from '@public/i18next/i18n';

  // Create a locale specific timestamp
  $: time = new Date().toLocaleDateString($locale, {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
</script>

<main>
  <p>
    <select bind:value={$locale}>
      {#each locales as l}
        <option value={l}>{l}</option>
      {/each}
    </select>
  </p>

  <h1>{$t('title')}!</h1>
  <p>{@html $t('welcome', { name: 'Jane Doe' })}!</p>
  <p>{@html $t('time', { time })}!</p>
</main>

<style>
  main {
    max-width: 24rem;
    margin: 3rem auto;
    padding: 0 1.5rem;
  }
</style>
