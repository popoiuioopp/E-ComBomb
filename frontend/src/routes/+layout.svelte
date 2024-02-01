<script>
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<Navbar />

<slot />

<style>
	:global(body) {
		margin: 0;
		padding: 0;
	}
</style>
