<template>
  <v-container class="fill-height">
    <v-responsive class="align-center mx-auto" max-width="900">
      <div class="text-center">
        <h1 class="text-h2 font-weight-bold">🐙.☕️</h1>
        <h2 class="text-h5 font-weight-light">Password Checker</h2>
        <br />
        <v-btn icon :href="'https://github.com/nilgaar/PassCheck'" target="_blank">
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <br />
      </div>

      <div class="py-4"></div>
      <v-row>
        <v-col cols="12">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-col cols="12">
              <v-card class="py-2" color="surface-light" rounded="lg" variant="outlined">
                <div class="px-4">
                  <p class="text-caption">This tool is designed to enhance your security by checking if your password
                    has been exposed in any known breaches. By analyzing hashed versions of passwords against common
                    dictionaries and wordlists, it helps determine if your password is still secure.</p>
                </div>
              </v-card>
            </v-col>

            <v-form @submit.prevent="hashPassword">
              <h2 class="text-h5 font-weight-bold px-4">Enter a password</h2>
              <v-text-field v-model="password" label="Password" outlined type="password" class="px-4" required />
              <v-btn color="primary" class="mt-3 mx-4" type="submit">
                Check Password
              </v-btn>
            </v-form>
            <v-overlay opacity=".12" scrim="primary" :value="true" />
          </v-card>
        </v-col>
      </v-row>
      <!-- List of files -->
      <v-row>
        <v-col cols="12">
          <v-list>
            <v-subheader>Dictionaries containing the Password:</v-subheader>
            <v-list-item-group v-for="file in files" :key="file">
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>{{ file }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { sha3_512 } from "js-sha3";

const password = ref("");
const files = ref([]);

async function hashPassword() {
  const hashedPassword = sha3_512(password.value);
  console.log(hashedPassword);
  fetch("https://api.checker.pops.cafe?hash=" + hashedPassword)
    .then((response) => response.json())
    .then((data) => {
      files.value = data.data.split(', ').sort()
        .map((file: string) => file.trim());
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}
</script>
<style scoped>
/* Base text styling for white appearance */
.text-caption,
.text-h5,
p,
h1,
h2 {
  color: #FFFFFF;
  /* White text color */
  font-weight: 300;
  /* Lighter font weight for readability */
}

/* Responsive text styling */
.text-caption {
  padding: 0 20px;
  font-size: 1.2rem;
  /* Increased default font size */
}

@media (min-width: 600px) {
  .text-caption {
    font-size: 1.3rem;
    /* Slightly larger font size for tablets and small desktops */
  }
}

@media (min-width: 960px) {
  .text-caption {
    font-size: 1.5rem;
    /* Larger font size for larger screens */
  }
}

/* Adjustments for other text elements for consistency */
.text-h5,
h1,
h2 {
  font-size: 1.2rem;
  /* Adjust base sizes for headings */
}

@media (min-width: 600px) {
  .text-h5 {
    font-size: 1.4rem;
    /* Increase for medium screens */
  }

  h1 {
    font-size: 2.5rem;
    /* Larger for main title */
  }

  h2 {
    font-size: 2rem;
    /* Larger for secondary titles */
  }
}

@media (min-width: 960px) {
  .text-h5 {
    font-size: 1.6rem;
    /* Even larger for desktops */
  }

  h1 {
    font-size: 3rem;
    /* More prominent main title */
  }

  h2 {
    font-size: 2.5rem;
    /* More prominent secondary titles */
  }
}
</style>
