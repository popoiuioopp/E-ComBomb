<script script lang="ts">
	import { onMount } from 'svelte';

	import ProductCard from '$lib/components/ProductCard.svelte';
	import { ENDPOINTS } from '$lib/constants/endpoints';

	let allProducts: any[] = [];

	const getProducts = async () => {
		const response = await fetch(ENDPOINTS.getAllProducts, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});
		if (!response.ok) {
			return [];
		}

		const responseJson = await response.json();
		return responseJson.products;
	};

	onMount(async () => {
		allProducts = await getProducts();
	});
</script>

{#each allProducts as product}
	<ProductCard {product} />
{/each}
