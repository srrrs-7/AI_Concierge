<script lang="ts">
  export let contents: string[];

  let selected: string = '';
  let isOpen = false;

  function toggleDropdown() {
    isOpen = !isOpen;
  }

  function selectOption(option: any) {
    selected = option;
    toggleDropdown();
  }
</script>

<div class="dropdown" on:click={toggleDropdown} on:keydown={toggleDropdown}>
  <div class="selected-value">{selected}</div>
  {#if isOpen}
    <ul class="options">
      {#each contents as content}
        <li on:click={() => selectOption(content)} on:keydown={toggleDropdown}>
      {content}
        </li>
      {/each}
    </ul>
  {/if}
</div>

<style>
  .dropdown {
    position: relative;
    display: inline-block;
  }

  .selected-value {
    cursor: pointer;
    padding: 8px;
    border: 1px solid #ccc;
  }

  .options {
    position: absolute;
    top: 100%;
    left: 0;
    z-index: 1;
    list-style-type: none;
    padding: 0;
    margin: 0;
    background-color: #fff;
    border: 1px solid #ccc;
    max-height: 200px;
    overflow-y: auto;
  }

  .options li {
    padding: 8px;
    cursor: pointer;
  }

  .options li:hover {
    background-color: #f5f5f5;
  }
</style>
