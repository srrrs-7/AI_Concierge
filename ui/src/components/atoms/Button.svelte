<script lang="ts">
	export let btnName: string;
	export let fetchData: () => void;
</script>

<div>
	<button on:click={fetchData}>{btnName}</button>
</div>
