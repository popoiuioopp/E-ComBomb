<script lang="ts">
	import { goto } from '$app/navigation';
	import { PLACEHOLDER_IMAGE } from '$lib/constants/utils';

	export let order: {
		id: number;
		createdAt: Date;
		updatedAt: Date;
		deletedAt: Date | null;
		userId: number;
		items: Array<{
			id: number;
			orderId: number;
			productId: number;
			product: {
				id: number;
				name: string;
				price: number;
				description: string;
				productImage: string;
			};
			quantity: number;
		}>;
	};
</script>

<div class="order">
	<h2>Order ID: {order.id}</h2>
	<p>Order Date: {order.createdAt}</p>
	<button on:click={() => goto(`/order/${order.id}`)}>Order Details</button>
	{#each order.items as item (item.id)}
		<div class="item">
			<img
				src={item.product.productImage || PLACEHOLDER_IMAGE}
				alt={item.product.name}
				class="product-image"
			/>
			<div>
				<h3>{item.product.name}</h3>
				<p>{item.product.description}</p>
				<p>Quantity: {item.quantity}</p>
				<p>Price: {item.product.price}</p>
			</div>
		</div>
	{/each}
</div>

<style>
	.order {
		border: 1px solid #ccc;
		margin-bottom: 20px;
		padding: 10px;
	}

	.item {
		display: flex;
		margin-top: 10px;
	}

	.product-image {
		width: 100px;
		height: 100px;
		margin-right: 20px;
	}
</style>
