<script lang="ts">
	import Atom_Button from '../atoms/Button.svelte';

	function fetchData() {
		console.log('fetching data');
	}
</script>

<div>
	<Atom_Button btnName={'push'} {fetchData} />
</div>
