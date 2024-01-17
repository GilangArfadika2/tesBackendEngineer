<?php

namespace Database\Seeders;

use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Str;

class CustomerStoreSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        // generate data dummy di database Customer stores
        for ($i = 1; $i <= 10; $i++) {
            DB::table('customer_stores')->insert([
                'title' => 'Title ' . $i,
                'name' => 'Customer ' . $i,
                'gender' => ($i % 2 == 0) ? 'M' : 'W',
                'phone_number' => '021xxxxx' . $i,
                'image' => 'customer_image_' . $i . '.jpg',
                'email' => 'customer' . $i . '@example.com',
                'created_at' => now(),
            ]);
        }
    }
}