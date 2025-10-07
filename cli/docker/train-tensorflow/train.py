import os
import tensorflow as tf

# Load MNIST
(x_train, y_train), _ = tf.keras.datasets.mnist.load_data()
x_train = x_train / 255.0

# Simple model
model = tf.keras.Sequential([
    tf.keras.layers.Flatten(input_shape=(28,28)),
    tf.keras.layers.Dense(128, activation='relu'),
    tf.keras.layers.Dense(10)
])

model.compile(optimizer='adam',
              loss=tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True),
              metrics=['accuracy'])

# Train briefly
model.fit(x_train, y_train, epochs=1)

# Save model
os.makedirs("/workspace/output", exist_ok=True)
model.save("/workspace/output/model.h5")

print("✅ Training complete. Model saved to /workspace/output/model.h5")
