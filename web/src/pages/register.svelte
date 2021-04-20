<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  import { goto, url } from "@roxi/routify";
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
          username: username,
          email: email,
          password: password,
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
        error = new Error(err.message);
      }
    } catch (err) {
      console.log(err);
      error = new Error("something wrong happend :(");
    }
  }
</script>

<Navbar isLoggedIn={false}/>
<main>
  <form action="#" on:submit|preventDefault={register}>
    <h1>Register</h1>
    <div class="formcontainer">
      <hr />
      <div class="container">
        <label for="uname"><strong>Username</strong></label>
        <input
          type="text"
          placeholder="enter username"
          bind:value={username}
        />
        <label for="uname"><strong>Email</strong></label>
        <input
          type="text"
          placeholder="enter email"
          bind:value={email}
        />
        <label for="psw"><strong>Password</strong></label>
        <input
          type="password"
          placeholder="enter password"
          bind:value={password}
        />
        {#if error}
        <div class="error-container">
          <p style="color: #003049;">{error.message}</p>
        </div>
        {/if}
      </div>
      <button type="submit">Login</button>
      <div class="container" style="background-color: #eee">
        <span class="psw">
          <a href={$url("./forgot-password")}> Forgot password?</a>
        </span>
      </div>
    </div>
  </form>
</main>


<style>
  main {
    margin-left: 30%;
    margin-right: 30%;
  }
  form {
    border: 5px solid #f1f1f1;
  }

  input[type="text"],
  input[type="password"] {
    width: 100%;
    padding: 16px 8px;
    margin: 10px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 10px;
    box-sizing: border-box;
  }

  button {
    background-color: #003049;
    color: white;
    padding: 14px 0;
    margin: 10px 0;
    border: none;
    border-radius: 10px;
    width: 100%;
  }
  h1 {
    text-align: center;
    font-size: 18;
  }

  button:hover {
    opacity: 0.8;
  }

  .formcontainer {
    text-align: left;
    margin: 24px 50px 12px;
  }

  .container {
    padding: 16px 0;
    text-align: center;
  }

  .error-container {
    background-color: #c1121f;
    /* padding: 1px 0px; */
  }

  .error-container p {
    font-size: large;
  }
</style>
