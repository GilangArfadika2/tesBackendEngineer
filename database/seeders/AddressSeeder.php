<?php

namespace Database\Seeders;

use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;

class AddressSeeder extends Seeder
{
    /**
     * Run the database seeds.
     */
    public function run()
    {
        // generate data dummy di database address
        for ($i = 1; $i <= 10; $i++) {
            DB::table('address')->insert([
                'customer_store_id' => $i, 
                'address' => 'Address ' . $i,
                'district' => 'District ' . $i,
                'city' => 'City ' . $i,
                'province' => 'Province ' . $i,
                'postal_code' => 1,
                'created_at' => now(),
            ]);
        }
    }
}
