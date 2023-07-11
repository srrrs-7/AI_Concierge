<script lang="ts">
  let response: any;
  async function calendarHandler() {
    try {
      await fetch('api/get/calendar', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(res =>
        res.body
          ?.getReader()
          .read()
          .then(({ value} ) => {
            response = new TextDecoder().decode(value);
          })
      );
    } catch (error) {
      response = error;
    }
  }
</script>

{#if response}
  {response}
{/if}

<button on:click={calendarHandler}>
  calendar
</button>
