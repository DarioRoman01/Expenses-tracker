<script lang="ts">
  import { url, goto } from "@roxi/routify"
	export let isLoggedIn: boolean;
	
</script>


<nav class="nav">
  <div>
    <h2 class="navbar-title">Expenses Tracker</h2>
  </div>
  <div class="navbar-links">
		{#if isLoggedIn}
			<button class="navbar-expenses">
				Add Expenses
			</button>
			<button class="navbar-logout">
				logout
			</button>
		{:else}
			<button class="navbar-login" on:click={$goto("./login")}>
				Login
			</button>
			<button class="navbar-register" on:click={$goto("./register")}>
				Register
			</button>
		{/if}
  </div>
</nav>


<style>
  .nav {
    max-width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #003049;
    color: #ffffff;
    border-bottom: #ffffff 1px solid;
  }

  .navbar-links {
		display: flex;
    padding-top: 10px;
    align-content: center;
    text-decoration: none;
    font-size: large;
  }

  .navbar-login {
    background-color: #003049;
    border: #ffffff 1px solid;
    border-radius: 10px;
    color: #ffffff;
    margin-right: 5px;
    padding-left: 20px;
    padding-right: 20px;
  }
  
  .navbar-register {
    background-color: #003049;
    border: #ffffff 1px solid;
    border-radius: 10px;
    color: #ffffff;
    margin-right: 10px;
  }
  
  .navbar-title {
    padding-left: 10px;
  }

  .navbar-login:hover, .navbar-register:hover {
    background-color: #003b5a;
  }

	.navbar-expenses {
		background-color: #003049;
    border: #ffffff 1px solid;
    border-radius: 10px;
    color: #ffffff;
    margin-right: 10px;
	}

	.navbar-logout {
		background-color: #003049;
    border: #ffffff 1px solid;
    border-radius: 10px;
    color: #ffffff;
    margin-right: 5px;
    padding-left: 10px;
    padding-right: 10px;
	}

</style>