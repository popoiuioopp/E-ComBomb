<div class="product-card">
	<div class="product-image"></div>
	<h2 class="product-name">Test</h2>
	<p class="product-description">
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec efficitur augue nunc, in pretium
		ligula pharetra ac. Etiam auctor malesuada nulla vitae laoreet. Nam semper, velit accumsan
		malesuada tincidunt, ante nunc dictum neque, ut faucibus tellus risus sed erat. Integer nunc
		orci, dignissim sit amet ipsum nec, sagittis venenatis urna. Nam et consectetur orci. Nam augue
		lectus, consectetur non convallis in, suscipit non dolor. Donec ac pulvinar magna, nec tristique
		arcu. Donec leo ligula, ornare id est non, elementum tempus sapien. Pellentesque fermentum metus
		quis porttitor laoreet. Nunc viverra nec nunc nec vehicula. Aenean accumsan libero eu est
		ornare, sed dignissim mauris ultrices.
	</p>
	<p class="product-price">100 baht</p>
	<div class="product-actions">
		<button class="button-add-to-cart">Add to cart</button>
		<button class="button-buy-now">Buy Now</button>
	</div>
</div>

<style>
	.product-card {
		font-family: 'Arial', sans-serif;
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
	}

	.product-name {
		font-size: 1.5em;
		margin: 10px 0;
	}

	.product-description {
		color: #666;
		font-size: 0.9em;
		margin: 10px 0;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
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
</style>
