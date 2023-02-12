<template>
    <v-container>

        <!-- <h2>{{ team.name }}</h2> -->
        <h2>kahlys team</h2>

        <v-table>
            <thead>
                <tr>
                    <th class="text-left">Name</th>
                    <th class="text-left">Country</th>
                    <th class="text-left">Role</th>
                    <th class="text-left">Power</th>
                    <th class="text-left">Stamina</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in team.players" :key="item.id">
                    <td>{{ item.first_name }} {{ item.last_name }}</td>
                    <td>{{ item.country }}</td>
                    <td>{{ item.role }}</td>
                    <td>{{ item.power }}</td>
                    <td>{{ item.stamina }}</td>
                </tr>
            </tbody>
        </v-table>
    </v-container>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Team',

    data() {
        return {
            team: {},
        }
    },

    mounted() {
        axios
            .get("/api/team")
            .then(response => {
                if (response.status === 200) {
                    console.log(response)
                    this.team = response.data
                }
            }).catch(error => {
                console.log(error);
            });
    },
}
</script>
