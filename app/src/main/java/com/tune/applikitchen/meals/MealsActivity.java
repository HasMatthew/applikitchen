package com.tune.applikitchen.meals;

import android.app.Activity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import com.tune.applikitchen.R;


public class MealsActivity extends Activity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_meals);
        ListView mealsList = (ListView) findViewById(R.id.mealsList);
        mealsList.setOnScrollListener(new EndlessScrollListener() {
            @Override
            public void onLoadMore(int page, int totalItemsCount) {
                System.out.println("DERPDERP");
            }
        });

        String[] items = {
                "ONE",
                "TWO",
                "THREE!!!",
                "FOUR?",
                "ONE",
                "TWO",
                "THREE!!!",
                "FOUR?",
                "ONE",
                "TWO",
                "THREE!!!",
                "FOUR?",
                "ONE",
                "TWO",
                "THREE!!!",
                "FOUR?",
                "ONE",
                "TWO",
                "THREE!!!",
                "FOUR?"
        };
        ArrayAdapter<String> adapter = new ArrayAdapter<String>(this.getApplicationContext(), android.R.layout.simple_list_item_1, items);
        mealsList.setAdapter(adapter);
        adapter.notifyDataSetChanged();
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.meals, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();
        if (id == R.id.action_settings) {
            return true;
        }
        return super.onOptionsItemSelected(item);
    }
}
