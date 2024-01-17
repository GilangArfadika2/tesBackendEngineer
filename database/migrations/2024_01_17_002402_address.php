<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('address', function (Blueprint $table) {
        $table->id();
        $table->foreignId('customer_store_id')->constrained('customer_stores')->onDelete('cascade'); 
        $table->string('address');
        $table->string('district');
        $table->string('city');
        $table->string('province');
        $table->unsignedInteger('postal_code');
        $table->dateTime('created_at');
        $table->dateTime('updated_at')->nullable();
        $table->softDeletes(); 
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('address');
    }
};
