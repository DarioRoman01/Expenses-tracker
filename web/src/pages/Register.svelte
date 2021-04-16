<script lang="ts">
  import { goto } from "@roxi/routify";
  let username: string;
  let email: string;
  let password: string;
  let error: Error;

  async function register() {
    error = undefined;
    try {
      const res = await fetch("http://localhost:1323/signup", {
        method: "POST",
        body: JSON.stringify({
          username,
          email,
          password,
        }),
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (res.ok) {
        $goto("./home");
      } else {
        const err = await res.json();
        error = new Error(err);
      }
    } catch (err) {
      console.log(err);
      error = new Error("something wrong happend :(");
    }
  }
</script>

<div>
  <h1>Register</h1>
  <input type="text" bind:value={username} placeholder="Username" />
  <input type="email" bind:value={email} placeholder="Email" />
  <input type="password" bind:value={password} placeholder="Password" />
  {#if error}
    <p>{error.message}</p>
  {/if}
  <button on:click={register}>Register</button>
</div>
