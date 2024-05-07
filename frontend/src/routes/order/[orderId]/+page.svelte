<script lang="ts">
	import type { Order, OrderItem } from './+page';

	export let data;
	const order: Order = data.order;
	const getTotalPrice = (items: OrderItem[]): number => {
		let total: number = 0;
		for (let item of items) {
			total += item.product.price;
		}
		return total;
	};
</script>

<h1>Order Details</h1>

{#if order}
	<p><strong>Order ID:</strong> {order.id}</p>
	<p><strong>Status:</strong> {order.status}</p>
	<p><strong>Created At:</strong> {order.createdAt}</p>
	<p><strong>Updated At:</strong> {order.updatedAt}</p>
	<p><strong>User ID:</strong> {order.userId}</p>
	<p><strong>Total Price:</strong> ${getTotalPrice(order.items)}</p>

	<h2>Items in Order:</h2>

	<ul>
		{#each order.items as item (item.id)}
			<li>
				<p><strong>Product ID:</strong> {item.productId}</p>
				<p><strong>Product Name:</strong> {item.product.name}</p>
				<p><strong>Price:</strong> ${item.product.price}</p>
				<p><strong>Quantity:</strong> {item.quantity}</p>
				<p><strong>Description:</strong> {item.product.description}</p>
			</li>
		{/each}
	</ul>
{:else}
	<p>Loading order details...</p>
{/if}
