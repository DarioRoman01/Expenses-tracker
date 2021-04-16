<script lang="ts">
	import {url} from "@roxi/routify";
	import Navbar from "../components/Navbar.svelte"
</script>


<main>
	<Navbar />
	<h1>Welcome to expense tracker</h1>
	<a href={$url("./Register")}>Register</a>
	<a href={$url("./Login")}>Login</a>
</main>

