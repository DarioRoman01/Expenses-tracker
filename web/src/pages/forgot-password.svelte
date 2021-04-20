<script lang="ts">
  import Navbar from "../components/Navbar2.svelte";
  let email: string;
  let error: Error;
  let send: boolean;

  async function forgotPassword() {
    send = false;
    try {
      const res = await fetch("http://localhost:1323/forgot-password", {
        method: "post",
        body: JSON.stringify({
          email: email
        }),
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
          },
        credentials: "include",
      })
  
      if (res.ok) {
        console.log(res.json())
        send = true;
      } else {
        const err = await res.json();
        error = new Error(err.message)
      }
    }
    catch(err) {
      console.log(err)
      error = new Error("something wrong happend:(")
    }
  }
</script>

<Navbar isLoggedIn={false}/>
<main>
  {#if !send}
    <form action="#" on:submit|preventDefault={forgotPassword}>
      <h1>Forgot Password</h1>
      <div class="formcontainer">
        <label for="email"><strong>Email</strong></label>
        <input type="text" placeholder="enter email" bind:value={email} />
        {#if error}
          <div class="error-container">
            <p style="color: #003049;">{error.message}</p>
          </div>
        {/if}
      </div>
      <button type="submit">Send Email</button>
    </form>
  {:else}
    <h1>We send you and email</h1>
    <p class="success">Check your email to update your password</p>
  {/if}
</main>

<style>
  main {
    margin-left: 30%;
    margin-right: 30%;
  }

  form {
    border: 5px solid #f1f1f1;
  }

  input[type="text"] {
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

  .error-container {
    background-color: #c1121f;
    /* padding: 1px 0px; */
  }

  .success {
    text-align: center;
    font-size: 25px;
  }
</style>
