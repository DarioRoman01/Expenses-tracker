<script lang="ts">
	import {url} from "@roxi/routify";
	import Navbar2 from "../components/Navbar2.svelte"
	import { Card, CardText, CardActions, Button, MaterialApp } from 'svelte-materialify';
</script>


<main>
	<Navbar2 isLoggedIn={false}/>
	<h1>Welcome to expense tracker</h1>
	<a href={$url("./register")}>Register</a>
	<a href={$url("./login")}>Login</a>

	<MaterialApp>
		<div class="d-flex justify-center mt-4 mb-4">
			<Card outlined style="max-width:300px;">
				<div class="pl-4 pr-4 pt-3">
					<span class="text-overline"> OVERLINE </span>
					<br />
					<span class="text-h5 mb-2">Headline</span>
					<br />
				</div>
				<CardText>
					Lorem ipsum dolor sit amet consectetur adipisicing elit. Reprehenderit, qui quaerat
					rerum incidunt nisi ducimus?
				</CardText>
				<CardActions>
					<Button rounded outlined>Button</Button>
				</CardActions>
			</Card>
		</div>
		</MaterialApp>
</main>

