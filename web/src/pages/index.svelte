<script lang="ts">
	import {url} from "@roxi/routify";
	import Navbar2 from "../components/Navbar2.svelte"
	import { Card, CardText, CardActions, Button, MaterialApp } from 'svelte-materialify';
</script>


<main>
	<Navbar2 isLoggedIn={false}/>
	<div class="d-flex flex-column mt-6">
		<div class="d-flex justify-center">
				<h2 class="text-h2">
					Welcome to expenses tracker!
				</h2>
		</div>

		<div class="d-flex justify-center mt-3">
			<h3 class="text-h3">
				✨Track your expenses never be so easy✨
			</h3>
		</div>
	</div>
</main>

