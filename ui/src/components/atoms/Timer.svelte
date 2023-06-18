<script>
  import moment from 'moment';

  $: now = moment().format('HH:mm:ss');
  setInterval(() => {
    now = moment().format('HH:mm:ss');
  }, 1000);
</script>

<main>
  <h1>{now}</h1>
</main>

<style>
</style>
