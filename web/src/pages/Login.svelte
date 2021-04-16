<script lang="ts">
  import {goto} from "@roxi/routify"
  let usernameOrEmail: string;
  let password: string;
  let error: Error;

  async function login() {
    try {
      const res = await fetch("http://localhost:1323/register", {
        method: "POST",
        body: JSON.stringify({
          usernameOrEmail,
          password
        }),
        headers: {
          "Content-type": "application/json"
        }
      }) 
      if (res.ok) {
        $goto("./home")
      } else {
        error = new Error("there was an error :(")
      }
    } catch(err) {
      console.log(err)
      error = new Error("there was an error :(")
    }
  }

</script>

<main>
  <input type="text" bind:value={usernameOrEmail} placeholder="Username">
  <input type="password" bind:value={password} placeholder="Password">
  {#if error}
    <p>{error.message}</p>
  {/if}
  <button onclick={login}>Login</button>
</main>