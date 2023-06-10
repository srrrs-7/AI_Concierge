<script lang="ts">
	export let name: string = 'name';
</script>

<div>
	<h1>Hello {name}!</h1>
	<button on:click={() => console.log(name)}>Click me</button>
</div>
