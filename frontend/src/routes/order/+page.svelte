<script lang="ts">
	import { onMount } from 'svelte';
	import OrderCard from '$lib/components/OrderCard.svelte';
	import { ENDPOINTS } from '$lib/constants/endpoints';

	interface Product {
		id: number;
		name: string;
		price: number;
		description: string;
		userId: number;
		productImage: string;
	}

	interface OrderItem {
		id: number;
		orderId: number;
		productId: number;
		product: Product;
		quantity: number;
	}

	interface Order {
		id: number;
		createdAt: Date;
		updatedAt: Date;
		deletedAt: Date | null;
		userId: number;
		items: OrderItem[];
	}

	let orders: Order[] = [];
	let isLoading = true;

	onMount(async () => {
		isLoading = true;
		try {
			const response = await fetch(ENDPOINTS.getAllOrders, {
				method: 'GET',
				credentials: 'include',
				headers: {
					'Content-Type': 'application/json'
				}
			});

			if (!response.ok) {
				orders = [];
				throw new Error(`Failed to fetch orders:${response.statusText}`);
			}

			const data = await response.json();
			orders = data.orders;
			console.log(orders);
		} catch (error) {
			console.error('An error occurred while fetching orders:', error);
			orders = [];
		}
		isLoading = false;
	});
</script>

<h1>Orders</h1>

{#if orders.length === 0}
	<p>Your cart is empty.</p>
{:else}
	<div class="orders">
		{#each orders as order (order.id)}
			<OrderCard {order} />
		{/each}
	</div>
{/if}

<style>
	.orders {
		display: flex;
		flex-direction: column;
	}
</style>
