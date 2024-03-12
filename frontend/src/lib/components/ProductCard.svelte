<script lang="ts">
	export let product: {
		name: string;
		description: string;
		price: number;
		image: string | null;
	};

	let placeholderImage =
		'https://ichef.bbci.co.uk/news/976/cpsprodpb/16620/production/_91408619_55df76d5-2245-41c1-8031-07a4da3f313f.jpg';

	function addToCart() {

	}

	function buyNow() {

	}
</script>

<div class="product-card">
	<div class="product-image">
		<img src={product.image || placeholderImage} alt="product" />
	</div>
	<h2 class="product-name">{product.name}</h2>
	<p class="product-description">{product.description}</p>
	<p class="product-price">{product.price} baht</p>
	<div class="product-actions">
		<button on:click={addToCart} class="button-add-to-cart">Add to cart</button>
		<button on:click={buyNow} class="button-buy-now">Buy Now</button>
	</div>
</div>

<style>
	.product-card {
		width: 300px;
		border-radius: 20px;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		background-color: #fff;
		padding: 20px;
		box-sizing: border-box;
		border: 1px solid #ea722f;
		margin: 20px;
	}

	.product-image {
		background-color: #eaeaea;
		height: 160px;
		border-radius: 10px;
		margin-bottom: 20px;
		overflow: hidden;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.product-image img {
		max-height: 100%;
		max-width: 100%;
		object-fit: contain;
	}

	.product-name {
		font-size: 1.5em;
		margin: 10px 0;
	}

	.product-description {
		color: #666;
		font-size: 0.9em;
		margin: 10px 0;
		overflow: hidden;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		max-width: 200px;
	}

	.product-price {
		color: #333;
		font-size: 1.2em;
		margin: 20px 0;
		font-weight: bold;
	}

	.product-actions {
		display: flex;
		justify-content: space-around;
		margin-top: 20px;
	}

	.button-add-to-cart,
	.button-buy-now {
		padding: 10px 20px;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-weight: bold;
		text-transform: uppercase;
		transition: background-color 0.3s ease;
	}

	.button-add-to-cart {
		background-color: #fff;
		color: #ea722f;
		border: 2px solid #ea722f;
	}

	.button-add-to-cart:hover {
		background-color: #f5984c;
		color: #fff;
	}

	.button-buy-now {
		background-color: #ea722f;
		color: #fff;
	}

	.button-buy-now:hover {
		background-color: #f5984c;
	}

	.product-actions button:not(:last-child) {
		margin-right: 10px;
	}
</style>
