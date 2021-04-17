<script lang="ts">
	import {url} from "@roxi/routify";
	import Navbar from "../components/Navbar.svelte"
</script>


<main>
	<Navbar />
	<h1>Welcome to expense tracker</h1>
	<a href={$url("./register")}>Register</a>
	<a href={$url("./login")}>Login</a>
</main>

